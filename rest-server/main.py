from flask import Flask, json

companies = [{"id": 1, "name": "Company One"}, {"id": 2, "name": "Company Two"}]
company = {"id": 1, "name": "Company One"}

api = Flask(__name__)

import random

@api.route('/companies', methods=['GET'])
def get_companies():
  """Return list of companies from companies variable

  Fail in fail_rate% of cases with 503 error.
  """
  fail_rate = 70
  dice_roll = random.randint(0,100)
  if dice_roll < fail_rate:
    return "error", 503
  return json.dumps(companies)

@api.route('/company', methods=['GET'])
def get_company():
  """Return a single company from company variable

  Fail in fail_rate% of cases with 503 error.
  """
  fail_rate = 70
  dice_roll = random.randint(0,100)
  if dice_roll < fail_rate:
    return "error", 503
  return json.dumps(company)

if __name__ == '__main__':
    api.run()
