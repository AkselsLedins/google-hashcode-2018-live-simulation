<h1 align="center">Google Hash Code 2018 : live simulation</h1>


<div align="center">

[![CircleCI](https://circleci.com/gh/AkselsLedins/google-hashcode-2018-live-simulation.svg?style=shield)](https://circleci.com/gh/AkselsLedins/google-hashcode-2018-live-simulation)
![License](https://img.shields.io/github/license/AkselsLedins/google-hashcode-2018-live-simulation.svg)
![Forks](https://img.shields.io/github/forks/AkselsLedins/google-hashcode-2018-live-simulation.svg)
![Stars](https://img.shields.io/github/stars/AkselsLedins/google-hashcode-2018-live-simulation.svg)
![Twitter](https://img.shields.io/twitter/url/https/github.com/AkselsLedins/google-hashcode-2018-live-simulation.svg?style=social)

</div>

## Introduction

Some context on the competition:

We are given a list of pre-booked rides in a city and a fleet of self-driving vehicles. The objective of the competition is to assign the rides to vehicles, so
that riders get to their destinations on time.

[The competition subject](resources/subject.pdf)

#### Guide to better understanding the animation

| Illustration                                                                                                                                                        | Explanation                                                                         |
|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| ![driver](screenshots/legend/driver.png)                                                                                                                            | A red dot is a car.                                                                 |
| ![failed-trip](screenshots/legend/failed-trip.png)                                                                                                                  | The red line is a representation of a failed trip. The car arrived late.            |
| ![will-complete](screenshots/legend/driver-on-ride-1.png)  ![Preview](screenshots/legend/driver-on-ride-2.png)  ![Preview](screenshots/legend/driver-on-ride-3.png) | A blue line is a trip in progress. You can see the car going through.               |
| ![will-complete](screenshots/legend/driver-will-complete-1.png)  ![will-complete](screenshots/legend/driver-will-complete-2.png)                                    | A green line is a completed trip.                                                   |
| ![assigned](screenshots/legend/assigned-ride.png)                                                                                                                   | A grey line means that a driver is assigned to the trip.                            |
| ![waiting](screenshots/legend/driver-waiting-at-start.png)                                                                                                          | A yellow line means that the driver is ready, but it's too early to start the ride. |

**Important:** we are only showing the assigned trips. If you press [T] you can toggle off/on the unassigned ones.

<div>
<img alt="off" src="screenshots/legend/toggle-off-all-trips.png" width="45%" align="right" >
<img alt="on" src="screenshots/legend/toggle-on-all-trips.png" width="45%" >
</div>

## Preview

<div align="center">

![Preview](screenshots/preview.gif)

</div>

## Installation

```
$> make deps
$> make
```

## Running


| flag   | required | explanation                                                                               |
|--------|----------|-------------------------------------------------------------------------------------------|
| -o     | :white_check_mark:      | path to the output file you have generated with your program during the hashcode          |
| -i     | :white_check_mark:      | corresponding input file, ex: b_should_be_easy.in                                         |
| -noGui |          | you can run the simulation withtout a graphic interface. This will only output your score |
| -h     |          | display help                                                                              |

#### Example with a Graphic Interface

```
$> ./google-hash-code-2018 -o resources/output-files/b.out -i resources/input-files/b_should_be_easy.in
```


<div align="center">

![Preview](screenshots/v0.10.png)

</div>

#### Example without a Graphic Interface

```
$> go run main.go -o resources/output-files/e.out -i resources/input-files/e_high_bonus.in -noGui
16381105

```


## Graphic interface commands

| Command       | effect                                |
|---------------|---------------------------------------|
| press "space" | pause / start the simulation          |
| press "t"     | shows / hide all the simulation trips |
| arrow keys    | move the camera around                |
| mouse scroll  | zoom in/out                           |


## FAQ

### Why ?
I thought it would be a fun project to practice Go

## Contributing

See <a href="https://github.com/AkselsLedins/google-hashcode-2018-live-simulation/blob/master/CONTRIBUTING.md">Contributing.md</a>.

But basically feel free to contribute if you find something to improve.


## License
[GPL-V3](https://tldrlegal.com/license/gnu-general-public-license-v3-(gpl-3))
