# angle

## parsing

    c 10d dec               -- 10
    c 10° dec               -- 10
    c 10.5d dec             -- 10.5
    c 10d30m dec            -- 10.5
    c 10°30' dec            -- 10.5
    c 10°30′ dec            -- 10.5
    c [10° 30'] dec         -- 10.5
    c 10d30.3m dec          -- 10.505
    c 10d30m45s dec         -- 10.5125
    c 10°30′45″ dec         -- 10.5125
    c 10°30'45" dec         -- 10.5125
    c [10° 30′ 45″] dec     -- 10.5125
    c 10d30m45.18s dec      -- 10.51255
    c 10d30m45.18sW dec     -- -10.51255
    c [10d30m45.18s W] dec  -- -10.51255

    c 10.5d10.5m dms?       -- false
    c 10d-10.5m dms?        -- false
    c 10.5d10.5s dms?       -- false
    c 10.5m10.5s dms?       -- false
    c 10d20m30sfoo dms?     -- false
    c -10d20m30sW dms?      -- false

