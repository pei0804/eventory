## go-dtf

![travis-ci](https://travis-ci.org/yterajima/go-dtf.svg?branch=master)

go-dtf is a package for golang. go-dtf convert [W3C-DTF](http://www.w3.org/TR/NOTE-datetime) string to time.Time.

### Support Format

check: [Date and Time Formats](http://www.w3.org/TR/NOTE-datetime)

#### YYYY

ex. 2015

#### YYYY-MM

ex. 2015-12

#### YYYY-MM-DD

ex. 2015-12-10

#### YYYY-MM-DDThh:mmTZD

ex.
    - 2015-12-10T01:17Z
    - 2015-12-10T01:17+00:00
    - 2015-12-10T01:17+09:00

#### YYYY-MM-DDThh:mm:ssTZD

ex.
    - 2015-12-10T01:17:20Z
    - 2015-12-10T01:17:20+00:00
    - 2015-12-10T01:17:20+09:00

#### YYYY-MM-DDThh:mm:ss.sTZD

ex.
    - 2015-12-10T01:17:20.123456789Z
    - 2015-12-10T01:17:20.123456789+00:00
    - 2015-12-10T01:17:20.123456789+09:00
