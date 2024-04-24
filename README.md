# assessment-tax  โปรแกรมคำนวนภาษี

 assessment-tax เป็น Application คำนวนภาษี ที่ให้ผู้ใช้งานสามารถคำนวนภาษีบุคคลธรรมดา ตามขั้นบันใดภาษี พร้อมกับคำนวนค่าลดหย่อน และภาษีที่ต้องได้รับคืน

## Getting Started
 ```
 git clone https://github.com/LGROW101/assessment-tax.git
 
 cd assessment-tax
 
 docker compose up
```

## User stories

### Story: EXP01
`POST:` tax/calculations

```json
{
  "totalIncome": 500000.0,
  "wht": 0.0,
  "allowances": [
    {
      "allowanceType": "donation",
      "amount": 0.0
    }
  ]
}
```

Response body

```json
{
  "tax": 29000
}
```
----
### Story: EXP02

`POST:` tax/calculations

```json
{
  "totalIncome": 500000.0,
  "wht": 25000.0,
  "allowances": [
    {
      "allowanceType": "donation",
      "amount": 0.0
    }
  ]
}
```

Response body

```json
{
  "tax": 4000
}
```
----
### Story: EXP03

`POST:` tax/calculations

```json
{
  "totalIncome": 500000.0,
  "wht": 0.0,
  "allowances": [
    {
      "allowanceType": "donation",
      "amount": 200000.0
    }
  ]
}
```

Response body

```json
{
  "tax": 19000
}
```
----
### Story: EXP04

`POST:` tax/calculations

```json
{
  "totalIncome": 500000.0,
  "wht": 0.0,
  "allowances": [
    {
      "allowanceType": "donation",
      "amount": 200000.0
    }
  ],
  "IncludeTaxLevel": true
}
```

Response body

```json
{
  "tax": 19000,
  "taxLevel": [
    {
      "level": "0-150,000",
      "tax": 0.0
    },
    {
      "level": "150,001-500,000",
      "tax": 19000
    },
    {
      "level": "500,001-1,000,000",
      "tax": 0
    },
    {
      "level": "1,000,001-2,000,000",
      "tax": 0
    },
    {
      "level": "2,000,001 ขึ้นไป",
      "tax": 0
    }
  ]
}
```
----

### Story: EXP05

`POST:` /admin/deductions

```json
{
    "personalDeduction": 70000
} 
```

Response body

```json
{
    "personalDeduction": 70000
} 
```
----


### Story: EXP06

`POST:` tax/calculations/upload-csv
```
totalIncome,wht,donation
500000,0,0
600000,40000,20000
750000,50000,15000
```

Response body

```json
{
    "taxes": [
        {
            "tax": 29000,
            "totalIncome": 500000
        },
        {
            "tax": 0,
            "totalIncome": 600000
        },
        {
            "tax": 11250,
            "totalIncome": 750000
        }
    ]
}
```
-------
### Story: EXP07

`POST:` tax/calculations

```json
{
  "totalIncome": 500000.0,
  "wht": 0.0,
  "allowances": [
    {
      "allowanceType": "k-receipt",
      "amount": 200000.0
    },
    {
      "allowanceType": "donation",
      "amount": 100000.0
    }
  ],
  "IncludeTaxLevel": true
}
```

Response body

```json
{
    "tax": 14000,
    "taxLevel": [
        {
            "level": "0-150,000",
            "tax": 0
        },
        {
            "level": "150,001-500,000",
            "tax": 14000
        },
        {
            "level": "500,001-1,000,000",
            "tax": 0
        },
        {
            "level": "1,000,001-2,000,000",
            "tax": 0
        },
        {
            "level": "2,000,001 ขึ้นไป",
            "tax": 0
        }
    ]
}
```
----

### Story: EXP08

`POST:` /admin/deductions

```json
{
    "k_receipt": 70000
} 
```

Response body

```json
{
  "KReceipt": 70000
}
```
----
