#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <algorithm>
using namespace std;

int main(int argc, char** argv) {
    fstream file;
    file.open(argv[1]);
    vector<int> caloriesPerElf;

    string line;
    int totalCalories = 0;
    while(getline(file, line)) {
        if (line == "") {
            caloriesPerElf.push_back(totalCalories);
            totalCalories = 0;
        } else {
            totalCalories += stoi(line);
        }
    }

    sort(caloriesPerElf.begin(), caloriesPerElf.end(), std::greater<int>());
    cout << "First elf is carrying " << caloriesPerElf[0] << " calories." << endl;
    cout << "Total calories carried by elves #1-3: " << caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2] << endl;
}