s = gets.chomp

def isMatching?(s)
  opening = ["{", "(", "["]
  stuff = "qwertyuiopasdfghjklzxcvbnm'\":;<>?,./|\1234567890-=!@#%^&*~`".split(//)
  stack = []
  s.split(//).each_index do |i|
    if opening.include?(s[i])
      stack << [i, s[i]]
    elsif stuff.include?(s[i].downcase)
      next
    else
      if stack.empty?
        return i + 1
      end
      top = stack.pop
      if (top[1] == "{" && s[i] != "}") || (top[1] == "[" && s[i] != "]") || (top[1] == "(" && s[i] != ")")
        return i + 1
      end
    end
  end
  stack.empty? ? "Success" : stack.last[0] + 1
end

puts isMatching?(s)
