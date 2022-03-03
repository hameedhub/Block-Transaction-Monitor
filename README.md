# ERC1155
> ERC1155 Block transaction event listener for testnet or mainnet/ Dashboard API


## Installing / Getting started
Introduction on how to run the application.

> Clone the repo and cd into the file directory

```shell
go mod tidy --To install all dependencies 
go run main.go --To run the application
go build . -- To build the applicaiton
```

## Developing

Procedure for further development of this application:

```shell
git clone https://github.com/hameedhub/ERC1155
cd into the ERC1155 folder
cmd go mod tidy
```

## API Endpoints

The application endpoints include:
- `GET transactions?{filters}` ['sender','receiver', 'from_date', 'to_date']


## Configuration / API Endpoints

API endpoint only allows json format and the expected output are also in json format.

### Get Transaction
#### GET /transactions

On success response
```shell
{
    "status": true,
    "message": "success",
    "data": [
        {
            "id": 2645,
            "balance_sender": 9223372036.854776,
            "balance receiver": 0,
            "hash": "0x950755b905c4471fd24ff75dd095d8f73b9d563f6616bea7773ac11a6878e910",
            "sender": "0x30b8235f492265A734347C0bF36E2FFcAD887be2",
            "receiver": "0x1d6E8BAC6EA3730825bde4B005ed7B2B39A2932d",
            "gas_price": 471190250348,
            "gas": 0,
            "value": "0",
            "nonce": 95351,
            "block_hash": "0xc14a71b29751a989533416d7733474cc8c5534d924a8848c88556e3cdfc3d2d6",
            "block_number": 14313812,
            "txn_index": 2,
            "chain_id": "1",
            "max_priority_fee_per_gas": "439279389026",
            "max_fee_per_gas": "471190250348",
            "created_at": "2022-03-03T12:46:23.388703+01:00"
        },
        {
            "id": 2643,
            "balance_sender": 9223372036.854776,
            "balance receiver": 0,
            "hash": "0xf31e3403f15a0b74de3b481e9e6c6e49d3b49012bfa69edcc87f689905d5b526",
            "sender": "0x30b8235f492265A734347C0bF36E2FFcAD887be2",
            "receiver": "0x1d6E8BAC6EA3730825bde4B005ed7B2B39A2932d",
            "gas_price": 31910861322,
            "gas": 0,
            "value": "0",
            "nonce": 95350,
            "block_hash": "0xc14a71b29751a989533416d7733474cc8c5534d924a8848c88556e3cdfc3d2d6",
            "block_number": 14313812,
            "txn_index": 0,
            "chain_id": "1",
            "max_priority_fee_per_gas": "0",
            "max_fee_per_gas": "31910861322",
            "created_at": "2022-03-03T12:46:19.673296+01:00"
        }
    ]
}
```
Error response
```shell
{
    "status": false,
    "message": "Invalid rows supplied",
    "error_code": 1,
    "Data": null
}
```
