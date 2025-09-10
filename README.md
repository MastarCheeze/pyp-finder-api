# CAIE Past Year Paper Finder API 🔍

🌐 Website: https://mastarcheeze.github.io/pyp-finder/ \
📚 GitHub repository of website: https://github.com/MastarCheeze/pyp-finder

## Synopsis

This API finds past CAIE exam papers (IGCSE, O Levels and A Levels) using their paper code. It finds and returns a link
to the exam paper PDF file.

I made this because I got tired of having to manually search for exam papers in a search engine and clicking
through a bunch of links before getting the PDF I want.

Exam papers are retrieved from [BestExamHelp](https://bestexamhelp.com/).

## API

The API is hosted using Google Cloud Run and is available publicly at
https://pyp-finder-server-417289630154.asia-east1.run.app.

To use the API, pass the fields `code` and `type` as a query string into the API URL (values passed are case
insensitive).

- `code` - The paper code, commonly found at the bottom of each page in a CAIE exam paper. (e.g. `9709/23/M/J/21`)
- `type` - The type of the paper. Valid values are `qp` for question papers, `ms` for mark schemes, `insert` for inserts
  and `pre` for pre-release material.

The response will return status code 200 if a paper is found successfully, and return status code 400 if not. The
response data is in JSON format. The fields returned are:

- `success` - `1` if a URL is found and `0` if not.
- `url` - A URL to the PDF file of the paper. This field only exists when `success` is `1`.
- `message` - The error message. This field only exists when `success` is `0`.

> [!note]
>
> If the paper code contains the paper type (e.g. `0509/13/INSERT/M/J/21`), the `type` field is not required. The API
> can deduce the paper type from the code. This coding style is used by CAIE only for inserts and pre-release material.

### Example usage

Query URL:

```url
https://pyp-finder-server-417289630154.asia-east1.run.app/?code=9702/13/M/J/21&type=qp
```

Response:

```json
{
  "success": 1,
  "url": "https://bestexamhelp.com/exam/cambridge-international-a-level/physics-9702/2021/9702_s21_qp_13.pdf"
}
```
