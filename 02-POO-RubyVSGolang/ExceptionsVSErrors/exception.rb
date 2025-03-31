def divide(a, b)
  raise "Cannot divide by zero" if b == 0
  a / b
rescue => e
  puts e.message
end