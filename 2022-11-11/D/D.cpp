// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=1347

// 10406 - Cutting tabletops
// Bever Lumber hires beavers to cut wood. The company has recently received a shippment of tabletops. Each tabletop is a convex polygon. However, in this hard economic times of cutting
// costs the company has ordered the tabletops from a not very respectable but cheap supplier. Some of the tabletops have the right
// shape but they are slightly too big. The beavers have to chomp
// of a strip of wood of a fixed width from each edge of the tabletop
// such that they get a tabletop of a similar shape but smaller. Your
// task is to find the area of the tabletop after beavers are done.

// Input
// Input consists of a number of cases each presented on a separate
// line. Each line consists of a sequence of numbers. The first number is d the width of the strip of wood to be cut off of each edge
// of the tabletop in centimeters. The next number n is an integer
// giving the number of vertices of the polygon. The next n pairs
// of numbers present xi and yi coordinates of polygon vertices for
// 1 ≤ i ≤ n given in clockwise order. A line containing only two zeroes terminate the input.
// d is much smaller than any of the sides of the polygon. The beavers cut the edges one after another
// and after each cut the number of vertices of the tabletop is the same.

// Output
// For each line of input produce one line of output containing one number to three decimal digits in the
// fraction giving the area of the tabletop after cutting.

// Sample Input
// 2 4 0 0 0 5 5 5 5 0
// 1 3 0 0 0 5 5 0
// 1 3 0 0 3 5.1961524 6 0
// 3 4 0 -10 -10 0 0 10 10 0
// 0 0

// Sample Output
// 1.000
// 1.257
// 2.785
// 66.294

// vector, rotate, translate, then cutPolygon

#include <algorithm>
#include <cmath>
#include <iomanip>
#include <iostream>
#include <limits>
#include <vector>

struct point {
    double x, y;
    
    point() {
        x = y = 0.0;
    }

    point(double x, double y) : x(x), y(y) {}

    bool operator < (point other) const {
        if (fabs(x - other.x) < std::numeric_limits<double>::epsilon()) {
            return x < other.x;
        }

        return y < other.y;
    }

    bool operator == (point other) const {
        return (fabs(x - other.x) < std::numeric_limits<double>::epsilon() && fabs(y - other.y) < std::numeric_limits<double>::epsilon());
    }
};

int main() {
    double d, area, perimeter, a, b, c, theta, w;
    unsigned n, i;
    std::vector<point> P;

    std::cout << std::setprecision(3) << std::fixed;

    while (std::cin >> d >> n && n) {
        P.resize(n);
        area = perimeter = 0;

        for (point & p : P) {
            std::cin >> p.x >> p.y;
        }

        P.push_back(P.front());

        for (i = 0; i < n; i++) {
            area += P[i].x * P[i + 1].y - P[i].y * P[i + 1].x;
            perimeter += hypot(P[i].x - P[i + 1].x, P[i].y - P[i + 1].y);
        }

        area = fabs(area) / 2.0;
        P[n + 1] = P[1];

        for (i = 1; i <= n; i++) {
            a = hypot(P[i].x - P[i - 1].x, P[i].y - P[i - 1].y);
            b = hypot(P[i].x - P[i + 1].x, P[i].y - P[i + 1].y);
            c = hypot(P[i + 1].x - P[i - 1].x, P[i + 1].y - P[i - 1].y);
            theta = acos((a * a + b * b - c * c ) / (2.0 * a * b)) / 2.0;
            w = d / tan(theta);
            perimeter -= w + w;
            area -= d * w;
        }

        area -= perimeter * d;

        std::cout << area << std::endl;
    }

	return 0;
}
