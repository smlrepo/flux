from(bucket:"test")
    |> range(start: 2018-11-07T00:00:00Z)
    |> highestAverage(n: 3, by: ["_measurement", "host"])
