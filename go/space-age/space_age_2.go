package space

type Planet string

const EARTH_YEAR = 31557600.0

var planetMap = map[Planet]float64 {
    "Earth": EARTH_YEAR,
    "Mercury": 0.2408467*EARTH_YEAR,
    "Venus": 0.61519726*EARTH_YEAR,
    "Mars": 1.8808158*EARTH_YEAR,
    "Jupiter": 11.862615*EARTH_YEAR,
    "Saturn": 29.447498*EARTH_YEAR,
    "Uranus": 84.016846*EARTH_YEAR,
    "Neptune": 164.79132*EARTH_YEAR,
}

func Age(seconds float64, planet Planet) float64 {

    return seconds / planetMap[planet]
}
