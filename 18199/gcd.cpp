/*
 * gcd.cpp
 *
 *  Michael Novak
 *  I pledge my honor that I have abided by the stevens honor system.
 */

#include <iostream>
#include <sstream>
using namespace std;

int gcd_recursive(int m, int n)
{
    if (n != 0)
       return gcd_recursive(n, m % n);
    else
       return abs(m);
}

int gcd_iterative(int m, int n)
{
    while ( n != 0) {
        long r = m % n;
        m = n;
        n = r;
    }
    return abs(m);
}

int main(int argc, char *argv[]) {
	int m, n;
	istringstream iss;

	if (argc != 3) {
		cerr << "Usage: " << argv[0] << " <integer m> <integer n>"<< endl;
		return 1;
	}
	iss.str(argv[1]);
	if ( !(iss >> m) ) {
		cerr << "Error: The first argument is not a valid integer."<< endl;
		return 1;
	}
	iss.clear(); // clear the error code
	iss.str(argv[2]);
	if ( !(iss >> n) ) {
		cerr << "Error: The second argument is not a valid integer."<< endl;
		return 1;
	}
	cout << "Iterative: gcd(" << m << ", " << n << ") = " << gcd_iterative(m, n) << endl;
	cout << "Recursive: gcd(" << m << ", " << n << ") = " << gcd_recursive(m, n) << endl;
	return 0;
}
