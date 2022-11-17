=begin
Write your code for the 'Sum Of Multiples' exercise in this file. Make the tests in
`sum_of_multiples_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/sum-of-multiples` directory.
=end

class SumOfMultiples
  def initialize(*args)
    @multiples = *args
  end

  def to(limit)
    count = 0
    list = []
    if @multiples.empty? || @multiples.first == 0
      return 0
    end
    while count < limit do
      @multiples.each do |m|
        if m == 0
          next
        elsif count % m == 0
          list << count
          break
        end 
      end
      count += 1
    end
    list.inject(:+)
  end
end
