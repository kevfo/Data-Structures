pattern = gets.chomp
string = gets.chomp

def polyhash(string, p, x)
  hash = 0
  (string.length - 1).downto(0).each do |i|
    hash = ((hash * x) + string[i].ord) % p
  end
  hash
end

def precompute_hashes(string, pattern, p, x)
  h = Array.new(string.length - pattern.length + 1, 0)
  s = string[(string.length - pattern.length)...string.length]
  h[string.length - pattern.length] = polyhash(s, p, x)
  y = 1

  1.upto(pattern.length).each do |i|
    y = (y * x) % p
  end

  (string.length - pattern.length - 1).downto(0).each do |i|
    h[i] = ((x * h[i + 1]) + string[i].ord - (y * string[i + pattern.length].ord)) % p
  end
  h
end

def naive_rabin(string, pattern)
  p = 6700417
  x = rand(1..p)
  result = []
  pHash = polyhash(pattern, p, x)

  0.upto(string.length - pattern.length).each do |i|
    tHash = polyhash(string[i...(i + pattern.length)], p, x)
    if tHash != pHash
      next
    end
    if string[i...(pattern.length + i)] == pattern
      result << i
    end
  end
  result
end

def better_rabin(string, pattern)
  p = 6700417
  x = rand(1..p)
  result = []
  pHash = polyhash(pattern, p, x)
  h = precompute_hashes(string, pattern, p, x)

  0.upto(string.length - pattern.length).each do |i|
    if pHash != h[i]
      next
    end
    if string[i...(i + pattern.length)] == pattern
      result << i
    end
  end
  result
end

puts better_rabin(string, pattern)
