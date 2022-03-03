bucketsize = gets.chomp.to_i
numOperations = gets.chomp.to_i

dict = {} ; commands = []
numOperations.times do |i|
  command, input = gets.chomp.split(/ /)
  commands << [command, input]
end

def polyhash(string, p, x, bucketsize)
  ans = 0
  0.upto(string.length - 1).each do |i|
    ans += string[i].ord * x**i
  end
  (ans % p) % bucketsize
end

commands.each do |i|
  key = polyhash(i[1], 1000000007, 263, bucketsize)

  case i[0]
  when 'add'
    dict[key] = dict.has_key?(key) ? dict[key] << i[1] : [i[1]]
  when 'del'
    dict.has_key?(key) ? dict[key].delete(i[1]) : next
  when 'find'
    puts dict.has_key?(key) ? "yes" : "no"
  when 'check'
    puts dict.has_key?(i[1].to_i) ? dict[i[1].to_i].reverse.join(" ") : " "
  end
end
