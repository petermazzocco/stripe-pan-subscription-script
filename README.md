# Stripe Bulk Subscription Migration Tool

A Go utility for creating bulk subscriptions for Stripe customers that have been transferred via Copy PAN data.


## Problem Solved

When migrating customers using Stripe's Copy PAN data feature, customer data is transferred but subscriptions are not. This Go tool bridges that gap by creating new subscriptions for transferred customers while preserving their original billing dates. You can read more about why Stripe doesn't copy subscription's over here: [https://docs.stripe.com/get-started/data-migrations/pan-copy-self-serve?copy-method=full&edit=true](https://docs.stripe.com/get-started/data-migrations/pan-copy-self-serve?copy-method=full&edit=true)

## Disclaimer

**USE AT YOUR OWN RISK**. The author is not responsible for any loss of funds, subscriptions, or other issues that may arise from using this tool. Any command you run is at your own risk. Always test thoroughly in a test environment before using in production.

## Features

- Reads customer data from CSV files
- Supports multiple date formats for purchase/subscription dates
- Obtains customer ID's for a list of emails given
- Calculates appropriate next billing dates
- Creates Stripe subscriptions with original billing cycle anchors
- Handles edge cases like leap years and varying month lengths
- Intelligently schedules billing to either same month (if current date < billing date) or next month (if current date > billing date)

## Prerequisites

- Go 1.22+
- Stripe API keys
- CSV file with customer data including purchase date, email, and customer ID (check below to get customer ID's)

## Installation

1. Clone this repository
2. Install dependencies:
   ```
   go mod download
   ```

## Configuration

Create a `.env` file in the project root with your Stripe API keys and price ID:

```
TEST_STRIPE_KEY=sk_test_...
PROD_STRIPE_KEY=sk_live_...
PRICE_ID=price_123...
```
To find your priceID, please go to your Stripe dashboard.

## CSV Format

Prepare a CSV file with the following columns:
1. Purchase Date - Original subscription/purchase date
2. Email - Customer email
3. Customer ID - Stripe customer ID (cus_...)
Add any additional fields here

Example:
```
Purchase Date,Email,Name,Plan,Customer ID
4/30/25 9:33,customer@example.com,John Doe,Premium,cus_123456789
```

## Usage

The project uses a Makefile to simplify build and run operations for a test and production environment.
**CAUTION**: Make sure you run `make run-test` before running any operations in a production environment.

### Building the project

Build both environments:
```
make all
```

Build test environment only:
```
make build-test
```

Build production environment only:
```
make build-prod
```

### Running the project

Run test environment:
```
make run-test
```

Run production environment:
```
make run-prod
```

Obtain customer ID's from stripe via emails:
```
make run-customer
```

Clean build artifacts:
```
make clean
```

## Output

The script outputs details about each created subscription, including:
- Customer email
- Original purchase date
- Next calculated billing date
- Subscription ID

## Error Handling

The script includes graceful error handling for:
- Invalid date formats
- Missing customer data
- Stripe API errors


## License

MIT License

Copyright (c) 2025

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sub-license, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
