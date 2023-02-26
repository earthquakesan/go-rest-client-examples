import logging
import sys
from flask import Flask, json, request

companies = [{"id": 1, "name": "Company One"}, {"id": 2, "name": "Company Two"}]
company = {"id": 1, "name": "Company One"}

api = Flask(__name__)

import random

@api.route('/companies', methods=['POST'])
def post_companies():
  """Save list of companies

  Fail in fail_rate% of cases with 503 error.
  """
  print(request.json)
  if simulate_failure():
    return "error", 503
  return "", 204

@api.route('/companies', methods=['GET'])
def get_companies():
  """Return list of companies from companies variable

  Fail in fail_rate% of cases with 503 error.
  """
  if simulate_failure():
    return "error", 503
  return json.dumps(companies)

@api.route('/company', methods=['PUT'])
def put_company():
  """Save a single company

  Won't actually do anything, just output the body of PUT request
  Fail in fail_rate% of cases with 503 error.
  """
  # Client must set Content-Type: application/json header
  # See: https://github.com/pallets/werkzeug/blob/main/src/werkzeug/wrappers/request.py#L572
  print(request.json)
  if simulate_failure():
    return "error", 503
  return "", 204

@api.route('/company', methods=['GET'])
def get_company():
  """Return a single company from company variable

  Fail in fail_rate% of cases with 503 error.
  """
  if simulate_failure():
    return "error", 503
  return json.dumps(company)

def simulate_failure():
  fail_rate = 50
  dice_roll = random.randint(0,100)
  if dice_roll < fail_rate:
    return True
  return False

if __name__ == '__main__':
    api.run()
