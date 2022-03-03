numCommands = gets.chomp.to_i

commands = []
numCommands.times do |i|
  command, number, person = gets.chomp.split(/ /)
  commands << [command, number, person]
end

def processCommands(commandList)
  phoneBook = {}

  commandList.each do |i|
    if i[0] == "add"
      phoneBook[i[1]] = i[2]
    elsif i[0] == "find"
      puts phoneBook.has_key?(i[1]) ? phoneBook[i[1]] : "Not found"
    elsif i[0] == "del"
      phoneBook.delete(i[1])
    end
  end
end

processCommands(commands)
