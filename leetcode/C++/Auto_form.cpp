#include<iostream>
#include<fstream>

using namespace std;

int main(int argc, char* argv[])
{
	char buffer[110];
	ofstream outfile;
	ifstream infile;
	
	if(argc < 1)
	{
		cout << "file is error!" << endl; 
		return -1;
	}
	outfile.open("form.cpp",ofstream::out | ofstream::app);
	infile.open(argv[1], ifstream::in);
	
	if(!infile.is_open())
	{
		cout <<"Error opening file"<< endl;
	}

	int count = 0;
	while(!infile.eof())
	{
		infile.getline(buffer,100);
		if((count % 2) == 0)
			outfile << buffer << endl;
		count++;
	}

	outfile.close();
	infile.close();
}
