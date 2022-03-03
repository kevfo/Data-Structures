string = gets.chomp
numRanges = gets.chomp.to_i

ranges = []
numRanges.times do |i|
  range = gets.chomp.split(/ /).map(&:to_i)
  ranges.push(range)
end

def makeHash(string, m)
  h = Array.new(string.length, 0)
  x = rand(1..10**9)
  (1...string.length).each do |i|
    h[i] = ((x * h[i - 1]) + string[i].ord) % m
  end
  h
end

def polyHash(string, p, x)
  hash = 0
  (string.length - 1).downto(0).each do |i|
    hash = ((hash * x) + string[i].ord) % p
  end
  hash
end

def solve(string, ranges)
  m1, m2 = 10**9 + 7, 10**9 + 9
  h1, h2 = makeHash(string, m1), makeHash(string, m2)
  ranges.each do |i|
    if h1[i[0]...(i[0] + i[2])] == h1[i[1]...(i[1] + i[2])] && h2[i[0]...(i[0] + i[2])] == h2[i[1]...(i[1] + i[2])]
      puts "Yes"
    else
      puts "No"
    end
  end
end

solve(string, ranges)
