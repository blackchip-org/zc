# examples

<ul>
<li><a href='#seq'>seq</a></li>
<li><a href='#time'>time</a></li>
</ul>

## angle

Degrees, minutes, seconds:

<!-- test: dms -->

| Input                 | Stack
|-----------------------|-------------
| `-77.016389`          | `-77.016389`
| `dm`                  | `-77° 0.98334′`   
| `dms`                 | `fo`

## seq 

Even numbers:

<!-- test: even-numbers -->

| Input                 | Stack
|-----------------------|-------------
| `1 6 seq`             | `1 \| 2 \| 3 \| 4 \| 5 \| 6`
| `[2 mod 0 eq] filter` | `2 \| 4 \| 6`

Powers of two:

<!-- test: powers-of-two -->

| Input                 | Stack
|-----------------------|-------------
| `1 8 seq`             | `1 \| 2 \| 3 \| 4 \| 5 \| 6 \| 7 \| 8`
| `[2 swap pow] map`    | `2 \| 4 \| 8 \| 16 \| 32 \| 64 \| 128 \| 256`

## time 

Time worked:

<!-- test: time-worked -->

| Input           | Stack
|-----------------|-------------
| `6:14pm 9:49am` | `6:14pm \| 9:49am`
| `sub`           | `8h 25m`
| `minutes`       | `505` 

Days until Christmas (today is mocked to be 15 March):

<!-- test: days-until-christmas -->

| Input          | Stack
|----------------|-------------
| `3/15/24 now=` | *now set to 'Fri Mar 15 2024 12:00:00am -0400 EDT'* 
| `12/25 doy`    | `360`
| `now doy`      | `360 \| 75`
| `sub`          | `285`

