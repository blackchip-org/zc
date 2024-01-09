## time - examples

Time elapsed:

<!-- test: time-elapsed -->

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

