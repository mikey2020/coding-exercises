module SavingsAccount
  def self.interest_rate(balance)
    if balance < 0
      -3.213
    elsif balance < 1000
      0.5
    elsif balance >= 1000 && balance < 5000
      1.621
    elsif balance >= 5000
      2.475
    end
  end

  def self.annual_balance_update(balance)
    balance + (balance.to_f.abs * self.interest_rate(balance)/100)
  end

  def self.years_before_desired_balance(current_balance, desired_balance)
    ((desired_balance - current_balance)/(current_balance.to_f * self.interest_rate(current_balance)/100)).to_i
  end
end
