module Animal
  def make_sound
    raise NotImplementedError
  end
end

class Dog
  include Animal

  def make_sound
    "Guau"
  end
end