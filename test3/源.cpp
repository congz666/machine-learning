#include <iostream>
#include<vector>
#include<list>
#include <time.h>
#include<math.h>
#include<float.h>
using namespace std;

class Cluster
{
private:
	int SampleNum;				//������
	int ClusterNum;				//������
	int featurenum;				//ÿ������������
	int* ClusterResult;
	int MaxTimes;				//����������
	double** Sample;
	double** Distances;			//�������			
	double** Centers;			

public:
	void GetClusting(vector<std::vector<std::vector<double> > >& r, double** feateres, int ClusterNum, int SampleNum, int datanum);

private:
	void Init(double** feateres, int ClusterNum, int SampleNum, int datanum);				//��ĳ�ʼ��
	void k_means(vector<vector<vector<double> > >& r);											
	void k_means_Init();																	//�������ĵĳ�ʼ��
	void k_means_Calculate(vector<vector<vector<double> > >& r);
};

void Cluster::GetClusting(vector<std::vector<std::vector<double> > >& r, double** feateres, int ClusterNum, int SampleNum, int datanum)
{
	Init(feateres, ClusterNum, SampleNum, datanum);
	k_means(r);
}


/*�������ݳ�ʼ��*/
void Cluster::Init(double** feateres, int ClusterNum, int SampleNum, int datanum)
{
	Sample = feateres;
	featurenum = datanum;
	SampleNum = SampleNum;
	ClusterNum = ClusterNum;
	MaxTimes = 50;

	Centers = new double* [ClusterNum];
	for (int i = 0; i < ClusterNum; ++i)
	{
		Centers[i] = new double[featurenum];
	}

	Distances = new double* [SampleNum];
	for (int i = 0; i < SampleNum; ++i)
	{
		Distances[i] = new double[ClusterNum];
	}

	ClusterResult = new int[SampleNum];
}

void Cluster::k_means(vector<vector<vector<double> > >& r)
{
	k_means_Init();
	k_means_Calculate(r);
}


/*��ʼ����������*/
void Cluster::k_means_Init()
{
	for (int i = 0; i < ClusterNum; ++i)
	{
		for (int k = 0; k < featurenum; ++k)
		{
			Centers[i][k] = Sample[i][k];
		}
	}
}

/*�������*/
void Cluster::k_means_Calculate(vector<vector<vector<double> > >& r)
{

	double J = DBL_MAX;
	int time = MaxTimes;

	while (time)

	{
		double now_J = 0;
		--time;

		//�����ʼ��
		for (int i = 0; i < SampleNum; ++i)
		{
			for (int j = 0; j < ClusterNum; ++j)
			{
				Distances[i][j] = 0;

			}
		}

		for (int i = 0; i < SampleNum; ++i)
		{
			for (int j = 0; j < ClusterNum; ++j)
			{
				for (int k = 0; k < featurenum; ++k)
				{
					Distances[i][j] += abs(pow(Sample[i][k], 2) - pow(Centers[j][k], 2));
				}
				now_J += Distances[i][j];
			}
		}

		if (J - now_J < 0.01)
		{
			break;
		}
		J = now_J;

		//temp���ڴ����ʱ������
		vector<vector<vector<double> > > temp(ClusterNum);
		for (int i = 0; i < SampleNum; ++i)
		{

			double min = DBL_MAX;
			for (int j = 0; j < ClusterNum; ++j)
			{
				if (Distances[i][j] < min)
				{
					min = Distances[i][j];
					ClusterResult[i] = j;
				}
			}

			vector<double> vec(featurenum);
			for (int k = 0; k < featurenum; ++k)
			{
				vec[k] = Sample[i][k];
			}
			temp[ClusterResult[i]].push_back(vec);
			
		}
		r = temp;

		for (int j = 0; j < ClusterNum; ++j)
		{
			for (int k = 0; k < featurenum; ++k)
			{

				Centers[j][k] = 0;
			}
		}

		for (int j = 0; j < ClusterNum; ++j)
		{
			for (int k = 0; k < featurenum; ++k)
			{
				for (int s = 0; s < r[j].size(); ++s)
				{
					Centers[j][k] += r[j][s][k];
				}
				if (r[j].size() != 0)
				{
					Centers[j][k] /= r[j].size();
				}
			}
		}
	}

	//�����������
	for (int j = 0; j < ClusterNum; ++j)
	{
		for (int k = 0; k < featurenum; ++k)
		{
			cout << Centers[j][k] << " ";
		}
		cout << endl;
	}
}

//groupnum ��������  
//datanum ÿ���麬�е����ݸ���
double** createdata(int groupnum, int datanum)
{
	srand((int)time(0));
	double** data = new  double* [groupnum];
	for (int i = 0; i < groupnum; ++i)
	{
		data[i] = new double[datanum];
	}
	cout << "�������ݣ�" << endl;
	for (int i = 0; i < groupnum; ++i)
	{
		for (int j = 0; j < datanum; ++j)
		{
			data[i][j] = ((int)rand() % 40001-20000)/10000.0;
			cout << data[i][j] << "	  ";
			if(j % 10 == 9)
			{
				cout << endl;
			}
		}
		cout << endl;
	}
	return data;
}

int main()
{
	vector<std::vector<std::vector<double> > >r;
	Cluster exam;
	double** data = createdata(10, 100);
	exam.GetClusting(r, data, 5, 10, 100);
	for (int i = 0; i < r.size(); ++i)
	{
		cout << endl;
		cout << "��" << i + 1 << "��" << endl;
		for (int j = 0; j < r[i].size(); ++j)
		{
			for (int k = 0; k < r[i][j].size(); ++k)
			{
				cout << r[i][j][k] << "   ";
			}
			cout << endl;
		}
	}
}
