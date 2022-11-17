// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=1238

// 10297 - Beavergnaw
// When chomping a tree the beaver cuts a very specific shape out of the tree trunk. What is left in
// the tree trunk looks like two frustums of a cone joined by a cylinder with the diameter the same as
// its height. A very curious beaver tries not to demolish a tree but rather sort out what should be the
// diameter of the cylinder joining the frustums such that he chomped out certain amount of wood. You
// are to help him to do the calculations.
// We will consider an idealized beaver chomping an idealized tree. Let us assume that the tree trunk
// is a cylinder of diameter D and that the beaver chomps on a segment of the trunk also of height D.
// What should be the diameter d of the inner cylinder such that the beaver chmped out V cubic units of
// wood?

// Input
// Input contains multiple cases each presented on a separate line. Each line contains two integer numbers
// D and V separated by whitespace. D is the linear units and V is in cubic units. V will not exceed the
// maximum volume of wood that the beaver can chomp. A line with D = 0 and V = 0 follows the last
// case.

// Output
// For each case, one line of output should be produced containing one number rounded to three fractional
// digits giving the value of d measured in linear units.

// Sample Input
// 10 250
// 20 2500
// 25 7000
// 50 50000
// 0 0

// Sample Output
// 8.054
// 14.775
// 13.115
// 30.901

// cones, cylinders, volumes

#include <cmath>
#include <iomanip>
#include <iostream>

int main() {
    unsigned D, V;
    std::cout << std::setprecision(3) << std::fixed;

    while (std::cin >> D >> V && D != 0 && V != 0) {
        std::cout << cbrt(D*D*D - 6.0*V/M_PI) << std::endl;
    }

	return 0;
}
