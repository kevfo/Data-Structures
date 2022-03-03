numCommands = gets.chomp.to_i

commands = []
numCommands.times do |i|
  command = gets.chomp.split(/ /)
  commands << command
end

def processCommands(commandList)
  stack, maxes = [], []
  commandList.each do |i|
    if i[0] == "push"
      if maxes.empty? || i[1].to_i > maxes.last then maxes.push(i[1].to_i) end
      stack.push(i[1].to_i)
    elsif i[0] == "pop"
      if i[1].to_i == maxes.last then maxes.pop end
      stack.pop
    elsif i[0] == "max"
      puts maxes.last
    end
  end
end

processCommands(commands)
