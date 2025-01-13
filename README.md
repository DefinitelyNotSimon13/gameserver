Server
======

SGCP Version 0
--------------

-	Byte 0 | Version
-	Byte 1 | Type
-	Byte 2 - 17 | UserId
-	Byte 18 | Flags
-	Byte 19 - 22 | Payload Length
-	Byte 23 - X | Payload
-	(Byte X - Y | CRC Checksum)

### Flags

| Bit | Usage to Server       | Usage to Client | Notes |
|-----|-----------------------|-----------------|-------|
| 0   | Type (e.g. Init/Join) | Success         |       |
| 1   |                       |                 |       |
| 2   |                       |                 |       |
| 3   |                       |                 |       |
| 4   |                       |                 |       |
| 5   |                       |                 |       |
| 6   |                       |                 |       |
| 7   |                       |                 |       |
| 8   |                       |                 |       |

to client: if bit0 == 1 -> failure

### Types

#### TCP

##### To Server

| Int | Bits | Type                 | Payload           | Payload Len | Flags           |
|-----|------|----------------------|-------------------|-------------|-----------------|
| 0   |      | PlayerInit           | Username          | 32 Byte     |                 |
| 1   |      | SessionInit/Join REQ | Nothing/SessionId | 0/16Byte    | Bit 0 Init/Join |
| 2   |      |                      |                   |             |                 |
| 3   |      |                      |                   |             |                 |
| 4   |      |                      |                   |             |                 |
| 5   |      |                      |                   |             |                 |

##### To Client

| Int | Bits | Type                 | Payload           | Payload Len | Flags         |
|-----|------|----------------------|-------------------|-------------|---------------|
| 0   |      | PlayerInit ACK       | \-                | 0           | Bit 0 Success |
| 1   |      | SessionInit/Join ACK | Nothing/SessionId | 0/16 Byte   | Bit 0 Success |
| 2   |      |                      |                   |             |               |
| 3   |      |                      |                   |             |               |
| 4   |      |                      |                   |             |               |
| 5   |      |                      |                   |             |               |

#### UDP

| Int | Bits | Type          | Payload  | Payload Len | Flags |
|-----|------|---------------|----------|-------------|-------|
| 6   |      | UDP Discovery |          |             |       |
| 7   |      | UDP Message   | Position | ?           |       |

#### Shared

| Int | Bits     | Type  | Payload | Payload Len | Flags |
|-----|----------|-------|---------|-------------|-------|
| 255 | 11111111 | Error |         |             |       |
