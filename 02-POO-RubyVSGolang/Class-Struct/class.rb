class Person
  attr_accessor :name

  def initialize(name)
    @name = name
  end

  def saludar
    "Hi, i am #{@name}"
  end
end