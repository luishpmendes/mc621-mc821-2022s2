// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=799

// 858 - Berry Picking
// It is not easy for a bird to find food when winter approaches. The fields become empty and all the
// crops have been harvested. There are only some wild berries that can be picked, but their location is
// sometimes hidden.
// We will help the birds find food. A bird looks for food by flying over a field and scanning it for the
// edible berries. Suppose that we are given the shape of a berry area in the field and we want to decide
// whether it is useful for the bird to cross the field along a chosen line. The flight is considered useful if
// the bird flies over an extension of berries that exceeds some threshold length.
// We will view a berry area as a polygon where vertices are approximated by integer values. Birds
// always follow vertical lines in their flight, and never fly over the vertices of the polygon: in those points
// there are usually scarecrows that frighten them.
// Your task is to decide if the flight along a given line is useful. You will be given a sequence of
// coordinates defining the berry area, a value for the threshold and the x-coordinate of the line to be
// followed by the bird.
// As an example, the following figure illustrates the
// location of a berry area in a field, where “line1” and
// “line2” are possible flight paths. If the threshold is
// set to 2 length units, line 2 will be an interesting flight
// path but line 1 will be not.
// Given the shape of the berry area, the threshold
// and the x-coordinate of the line to be followed by the
// bird, decide whether the line is useful. Assume that
// the given x-coordinate does not intersect the polygon
// at any vertex.

// Input
// The first line of the input contains the number T of
// test cases, followed by T input blocks.
// The first line of each inpt block consists of one
// integer N, the number of vertices of the polygon representing the berry area. Each of the following N input lines contains the X-Y integer coordinates of one
// vertex. The following line has an integer indicating
// the threshold for the decision. The last line of each
// input block has an integer for the x-coordinate to be
// evaluated. The set of vertices starts at an arbitrary
// point in the polygon and there are at most 10000 vertices in a polygon.

// Output
// For each input block output a single line, containing ‘YES’ if the chosen line is useful and ‘NO’ otherwise.

// Sample Input
// 2
// 7
// 4 1
// 2 2
// 5 3
// 2 5
// 2 3
// 0 1
// 2 0
// 2
// 3
// 7
// 4 1
// 2 2
// 5 3
// 2 5
// 2 3
// 0 1
// 2 0
// 2
// 3

// Sample Output
// YES
// YES

// ver line-polygon intersect; sort; alternating segments

#include <algorithm>
#include <cmath>
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
    unsigned t, n, i, j;
    std::vector<point> polygon;
    double threshold, passx, pass, py;
    std::vector<double> passy;

    while (std::cin >> t) {
        while (t--) {
            std::cin >> n;

            polygon.resize(n);

            for (point & p : polygon) {
                std::cin >> p.x >> p.y;
            }

            std::cin >> threshold >> passx;

            pass = 0;
            passy.clear();

            for (i = 0, j = n-1; i < n; j = i++) {
                py = polygon[i].y + (polygon[j].y - polygon[i].y) * (passx - polygon[i].x) / (polygon[j].x - polygon[i].x);

                if ((polygon[i].x <= passx || polygon[j].x <= passx) && (polygon[i].x >= passx || polygon[j].x >= passx)) {
                    passy.push_back(py);
                }
            }

            std::sort(passy.begin(), passy.end());

            for (i = 0; i < passy.size(); i += 2) {
                pass += passy[i+1] - passy[i];
            }

            std::cout << (pass >= threshold ? "YES" : "NO") << std::endl;
        }
    }

	return 0;
}
