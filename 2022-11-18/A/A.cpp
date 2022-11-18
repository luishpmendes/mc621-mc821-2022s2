// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=226&page=show_problem&problem=2934

// 11834 - Elevator
// The FCC (Factory of Cylinders of Carbon) manufactures various types of cylinders of carbon. FCC
// is installed on the tenth floor of a building, and uses the several building’s elevators to transport the
// cylinders. For security, the cylinders must be transported in the upright position, and since they are
// heavy, at most two cylinders can be transported in a single elevator ride. The elevators have the shape
// of a parallelepiped and their height is always greater than the height of the cylinders.
// To minimize the number of elevator trips to transport the cylinders, the FCC wants, whenever
// possible, to put two cylinders in the elevator. The figure below illustrates, schematically (top view) a
// case where this is possible (a), and a case where this is not possible (b):
// As there is a very large amount of elevators and types of cylinders, FCC hired you to write a program
// that, given the dimensions of the elevator and of the two cylinders, determines whether it is possible to
// put the two cylinders in the elevator.

// Input
// The input contains several test cases. The first and only line of each test case contains four integers L,
// C, R1 and R2 , separated by blanks, indicating the width (1 ≤ L ≤ 100) and the length (1 ≤ C ≤ 100)
// of the elevator and the radii of the cylinders (1 ≤ R1, R2 ≤ 100).
// The last test case is followed by a line containing four zeros separated by blanks.

// Output
// For each test case your program should print a single line with a single character, ‘S’ if you can put the
// two cylinders in the elevator and ‘N’ otherwise.

// Sample Input
// 11 9 2 3
// 7 8 3 2
// 10 15 3 7
// 8 9 3 2
// 0 0 0 0

// Sample Output
// S
// N
// N
// S

// packing two circles in a rectangle

#include <cmath>
#include <iomanip>
#include <iostream>
#include <limits>

int main() {
    unsigned l, c, r1, r2;

    while (std::cin >> l >> c >> r1 >> r2 && (l || c || r1 || r2)) {
        if (2 * r1 > l || 2 * r1 > c || 2 * r2 > l || 2 * r2 > c
            || (r1 + r2 - l) * (r1 + r2 - l) + (r1 + r2 - c) * (r1 + r2 - c) < (r1 + r2) * (r1 + r2)) {
            std::cout << "N" << std::endl;
        } else {
            std::cout << "S" << std::endl;
        }
    }

	return 0;
}
