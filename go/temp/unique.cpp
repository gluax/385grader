/*******************************************************************************
 * Name        : unique.cpp
 * Author      : Michael Ficke II
 * Date        : 9/23/18
 * Description : Determining uniqueness of chars with int as bit vector.
 * Pledge      : I pledge my honor that I have abided by the Stevens Honor System -- Michael Ficke II
 ******************************************************************************/
#include <iostream>
#include <cctype>
#include <sstream>

using namespace std;

/**
 * given a string, we need to check if a string of characters is all lowercase or not.
 */

bool is_all_lowercase(const string &s) {

	int i = 0;
	char c;

	while (s[i]) {
		c = s[i];

		if (islower(c) == false) { //I found this technique on the c++ website for the function islower()
			return false;
		}

		i++;
	}
	return true;
}

bool all_unique_letters(const string &s) {
	// returns true if all letters in string are unique, that is
	// no duplicates are found; false otherwise.

	int length = s.length();
	unsigned int setter;
	unsigned int vector = 0; // we are going to use the setter and vector to check the bits of the char string to look for duplicates or not

	for (int i = 0; i < length; i++) {

		setter = 1 << (s[i] - 'a'); // 'a' represents 1 here in ASCII.  Since we are only dealing with characters here, this is okay to use for all strings given in this program

		//"and" check

		if (vector & setter)  //we found a duplicate
				{
			return false;
		}

		//"or" process, in other words, we did not find a duplicate yet and we need to keep looking

		vector = setter | vector;
	}

	return true;  //there are no duplicates

}

int main(int argc, char * const argv[]) {

//argc = 2, or at least it should be

	string input;

	istringstream iss; //we are going to use this to error check

	if (argc != 2) //we are creating an error message here
			{
		cerr << "Usage: " << argv[0] << " <string>" << endl;

		return 1;
	}

	iss.str(argv[1]);

	if (!(iss >> input)) //this is basically saying, if the input is not what we think it is (bad input)
	{
		cout << "Error: String must contain only lowercase letters." << endl;
		return 1;

	}

	if (is_all_lowercase(input) == false) {
		cout << "Error: String must contain only lowercase letters." << endl;
		return 1;

	}

	all_unique_letters(input);

	if (all_unique_letters(input)) { //messages for the function put down here for better programming technique
		cout << "All letters are unique.";
	} else {
		cout << "Duplicate letters found.";
	}

	return 0;

}
