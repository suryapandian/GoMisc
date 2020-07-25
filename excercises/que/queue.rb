class Stack
	def initialize()
		@mem = []
	end

	def push(i)
		@mem << i
	end

	def pop()
		@mem.pop
	end
end


class Queue
	def initialize()
		@mem = Stack.new()
	end

	def push(i)
		@mem.push(i)
	end

	def pop() //time complexity will increase in pop. we can move this code to push then, time complx will increase in push.
		temp = Stack.new
		swap(@mem,temp)
		res = temp.pop
		swap(temp,@mem)
		return res
	end

	def swap(one, two)
		i = one.pop
		while(i)
			two.push i
			i = one.pop
		end
	end
end

#Lok's implementation in Ruby