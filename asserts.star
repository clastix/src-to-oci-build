load("@ytt:data", "data")
load("@ytt:assert", "assert")

data.values.app_name != None or assert.fail("Missing app_name")
data.values.programming_language_runtime != None or assert.fail("Missing programming_language_runtime")
