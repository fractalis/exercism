package space

const (
    EARTH_SECONDS = 31557600.0
    MERCURY_RATIO = 0.2408467
    VENUS_RATIO = 0.61519726
    MARS_RATIO = 1.8808158
    JUPITER_RATIO = 11.862615
    SATURN_RATIO = 29.447498
    URANUS_RATIO = 84.016846
    NEPTUNE_RATIO = 164.79132
)

//type Planet string

func Age2(seconds float64, planet Planet) float64 {

    age := 0.0

    switch(planet) {
    case "Earth":
        age = (seconds/EARTH_SECONDS)
    case "Mercury":
        age = (seconds/(EARTH_SECONDS*MERCURY_RATIO))
    case "Venus":
        age = (seconds/(EARTH_SECONDS*VENUS_RATIO))
    case "Mars":
        age = (seconds/(EARTH_SECONDS*MARS_RATIO))
    case "Jupiter":
        age = (seconds/(EARTH_SECONDS*JUPITER_RATIO))
    case "Saturn":
        age = (seconds/(EARTH_SECONDS*SATURN_RATIO))
    case "Uranus":
        age = (seconds/(EARTH_SECONDS*URANUS_RATIO))
    case "Neptune":
        age = (seconds/(EARTH_SECONDS*NEPTUNE_RATIO))
    }

    return age
}
