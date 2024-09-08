An `Angle/DMS` value is an angle that can be expressed as:

- decimal degrees
- degrees and minutes
- degrees, minutes, and seconds

Any valid decimal number, such as 12.345 can be parsed as an Angle/DMS value.
Use unit markers to designate each part of the angle by using:

- degrees: `d`, `°`
- minutes: `m`, `'`, `′`
- seconds: `s`, `"`, `″`

Using the letter unit markers with no whitespace is the easiest to use when
entering values manually. All of the following parse to the same value:

<!-- test: ParseDMS -->

| Input                 | Stack
|-----------------------|-------------
| `c 10.5125 dec`       | `10.5125`
| `c 10.5125d dec`      | `10.5125`
| `c 10.5125° dec`      | `10.5125`
| `c 10d30.75m dec`     | `10.5125`
| `c 10d30.75' dec`     | `10.5125`
| `c [10° 30′ 45″] dec` | `10.5125`

