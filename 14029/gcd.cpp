/*
 * gcd.cpp
 *
 *  Created on: Aug 31, 2018
 *      Author: michael ficke II
 *
 *      I PLEDGE MY HONOR THAT I HAVE ABIDED BY THE STEVENS HONOR SYSTEM -- MICHAEL FICKE II
 */


#include<iostream>
#include<sstream> //need this for error checking
#include<cmath>  //need this for absolute value

using namespace std;

/**
 * the following code is the iterative gcd function
 */

int gcd_iterative(int m, int n)
{
int remainder = -1;  //never going to have a negative remainder.  This is a safe initial value for remainder

while(remainder != 0)
{
	if((m == 0) & (n == 0))  //gcd of zero is zero
	{
		return 0;
	}

	else if(m < 0)  //any negative values we need to take the absolute value of.  The math still makes sense
	{
		m = abs(m);
	}

	else if(n < 0)
	{
		n = abs(n);
	}

	else if((m == 0) | (n == 0))  //return the largest value if any of the numbers are zero, but not both zero
	{
		if(m > n)
		{
			return m;
		}

		else
		{
			return n;
		}
	}

	else if(m == n)  //if both numbers are equal, find GCD one way
	{
		remainder = m%n;
		m = n;
		n = remainder;
	}

	else if(m > n) //one number is larger than the other, we need to make sure we are taking the mod of the largest number first
	{
	remainder = m%n;
	m = n;
	n = remainder;

	}

else  //no need to try to mirror the code, just swap the m and n values so that m is larger than n
{
	int temp;  //we need to create a place holder variable so that we do not lose any information
	temp = m;
	m = n;
	n = temp;
}

}

return m;

}

//the following code is the recursive gcd function

int gcd_recursive(int m, int n)
{
if((m == 0) & (n == 0))  //if both of the numbers is zero, there's nothing else that we can do here
{
	return 0;
}

else if(m < 0)  //we need to get rid of any negative numbers
{
	return gcd_recursive(abs(m),n);
}

else if(n < 0)
{
	return gcd_recursive(m,abs(n));
}

else if((m == 0) | (n == 0))  //gcd of zero and a non-zero integer is just the non-zero integer
{
	if(m > n)
	{
		return m;
	}
	else
	{
		return n;
	}
}

else if(m == n)  //if the numbers are the same, then either number will function as the GCD
{
	return m;  //returning n would also be fine here as well
}

else  //actually find the GCD of the two numbers
{
if(m > n)
{
	return gcd_recursive(m-n,n);
}
else
{
	return gcd_recursive(m,n-m);
}

}

}

int main(int argc, char * argv[])
{
	int m, n;

	istringstream iss;  //a useful datatype that can convert the variable to other data types rather easily

	//we are creating error checking prompts here to make sure the inputs are correct

	if (argc != 3) //we are creating an error message here
	{
		cerr<<"Usage: "<<argv[0] << " <integer m> <integer n>" << endl; //check the notebook for what we are doing here with argv and argc.

		return 1;
	}

	iss.str(argv[1]);

	if(!(iss >> m) ) //these brackets are trying to convert the variable into the type m is and store it in m
	{
	cerr << "Error: The first argument is not a valid integer." << endl;
	return 1;

	}

	iss.clear(); //clears the error codes

	iss.str(argv[2]);

	if(!(iss >> n) ) //these brackets are trying to convert the variable into the type m is and store it in n
	{
	cerr << "Error: The second argument is not a valid integer." << endl;
	return 1;

	}

	cout<<"Iterative: gcd("<<m<<", "<<n<<") = "<<gcd_iterative(m,n)<<endl;
	cout<<"Recursive: gcd("<<m<<", "<<n<<") = "<<gcd_recursive(m,n)<<endl;

	return 0;
}

