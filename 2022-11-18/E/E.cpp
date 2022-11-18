// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=17&page=show_problem&problem=1518

// 10577 - Bounding box
// The Archeologists of the Current Millenium (ACM) now and then discover ancient artifacts located at
// vertices of regular polygons. The moving sand dunes of the desert render the excavations difficult and
// thus once three vertices of a polygon are discovered there is a need to cover the entire polygon with
// protective fabric.

// Input
// Input contains multiple cases. Each case describes one polygon. It starts with an integer n ≤ 50, the
// number of vertices in the polygon, followed by three pairs of real numbers giving the x and y coordinates
// of three vertices of the polygon. The numbers are separated by whitespace. The input ends with a n
// equal 0, this case should not be processed.

// Output
// For each line of input, output one line in the format shown below, giving the smallest area of a rectangle
// which can cover all the vertices of the polygon and whose sides are parallel to the x and y axes.

// Sample Input
// 4
// 10.00000 0.00000
// 0.00000 -10.00000
// -10.00000 0.00000
// 6
// 22.23086 0.42320
// -4.87328 11.92822
// 1.76914 27.57680
// 23
// 156.71567 -13.63236
// 139.03195 -22.04236
// 137.96925 -11.70517
// 0

// Sample Output
// Polygon 1: 400.000
// Polygon 2: 1056.172
// Polygon 3: 397.673

// get center+radius of outer circle from 3 points, get all vertices, get the min-x/max-x/min-y/max-y of the polygon

#include <cmath>
#include <iomanip>
#include <iostream>
#include <limits>
#include <vector>

struct point {
    double x, y;

    point() : x(0), y(0) {}

    point(double x, double y) : x(x), y(y) {}

    bool operator < (const point& p) const {
        if (fabs(x - p.x) > std::numeric_limits<double>::epsilon()) {
            return x < p.x;
        }

        return y < p.y;
    }

    bool operator == (const point& p) const {
        return fabs(x - p.x) < std::numeric_limits<double>::epsilon() && fabs(y - p.y) < std::numeric_limits<double>::epsilon();
    }
};

struct line {
    double a, b, c;
};

struct vec {
    double x, y;

    vec(double x, double y) : x(x), y(y) {}
};

double dist(point p1, point p2) {
    return hypot(p1.x - p2.x, p1.y - p2.y);
}

point rotate(point o, point p, double theta) {
    double dx = p.x - o.x,
           dy = p.y - o.y;
    return point(o.x + dx * cos(theta) - dy * sin(theta), o.y + dx * sin(theta) + dy * cos(theta));
}

line pointsToLine(point p1, point p2) {
    line l;

    if (fabs(p1.x - p2.x) < std::numeric_limits<double>::epsilon()) {
        l.a = 1.0;
        l.b = 0.0;
        l.c = -p1.x;
    } else {
        l.a = -(double)(p1.y - p2.y) / (p1.x - p2.x);
        l.b = 1.0;
        l.c = -(double)(l.a * p1.x) - p1.y;
    }

    return l;
}

bool areParallel(line l1, line l2) {
    return (fabs(l1.a-l2.a) < std::numeric_limits<double>::epsilon()) && (fabs(l1.b-l2.b) < std::numeric_limits<double>::epsilon());
}

bool areIntersect(line l1, line l2, point & p) {
    if (areParallel(l1, l2)) {
        return false;
    }

    p.x = (l1.b * l2.c - l2.b * l1.c) / (l1.a * l2.b - l2.a * l1.b);
    p.y = (l2.a * l1.c - l1.a * l2.c) / (l1.a * l2.b - l2.a * l1.b);

    return true;
}

vec toVec(point a, point b) {
    return vec(b.x - a.x, b.y - a.y);
}

vec scale(vec v, double s) {
    return vec(v.x * s, v.y * s);
}

point translate(point p, vec v) {
    return point(p.x + v.x, p.y + v.y);
}

double area(double ab, double bc, double ca) {
    double s = (ab + bc + ca) / 2.0;
    return sqrt(s * (s - ab) * (s - bc) * (s - ca));
}

double rCircumCircle(double ab, double bc, double ca) {
    return ab * bc * ca / (4.0 * area(ab, bc, ca));
}

double rCircumCircle(point a, point b, point c) {
    return rCircumCircle(dist(a, b), dist(b, c), dist(c, a));
}

line perpendicularBisectorFromLine(point a, point b, line l) {
    line result;
    point mid_point = point((a.x + b.x) / 2.0, (a.y + b.y) / 2.0);

    result.a = l.b;
    result.b = -l.a;
    result.c = -l.b * mid_point.x + l.a * mid_point.y;

    return result;
}

int circumCircle(point a, point b, point c, point & ctr, double & r) {
    r = rCircumCircle(a, b, c);

    if (fabs(r) < std::numeric_limits<double>::epsilon()) {
        return 0;
    }

    line l1, l2;

    l1 = pointsToLine(a, b);
    l2 = pointsToLine(b, c);

    l1 = perpendicularBisectorFromLine(a, b, l1);
    l2 = perpendicularBisectorFromLine(b, c, l2);

    if (!areIntersect(l1, l2, ctr)) {
        return 0;
    }

    return 1;
}

int main() {
    unsigned n, t, i;
    point a, b, c, ctr, p;
    double r, theta, min_x, max_x, min_y, max_y, ans;

    std::cout << std::setprecision(3) << std::fixed;

    for (t = 1; std::cin >> n && n; t++) {
        std::cin >> a.x >> a.y >> b.x >> b.y >> c.x >> c.y;
        ans = 0.0;
        theta = 2.0 * M_PI / n;

        if (circumCircle(a, b, c, ctr, r) != 0) {
            min_x = a.x;
            max_x = a.x;
            min_y = a.y;
            max_y = a.y;

            for (i = 1; i <= n; i++) {
                p = rotate(ctr, a, i * theta);
                min_x = std::min(min_x, p.x);
                max_x = std::max(max_x, p.x);
                min_y = std::min(min_y, p.y);
                max_y = std::max(max_y, p.y);
            }

            ans = (max_x - min_x) * (max_y - min_y);
        }

        std::cout << "Polygon " << t << ": " << ans << std::endl;
    }

	return 0;
}
