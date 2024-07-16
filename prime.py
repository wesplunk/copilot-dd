def is_prime(n):
    """
    Check if a number is prime.

    Parameters:
    n (int): The number to check.

    Returns:
    bool: True if the number is prime, False otherwise.
    """
    # Check if the number is less than 2, which is not prime
    if n < 2:
        return False

    # Check if the number is 2, which is prime
    if n == 2:
        return True

    # Check if the number is divisible by 2, which is not prime
    if n % 2 == 0:
        return False

    # Check for divisibility by odd numbers from 3 to the square root of n
    for i in range(3, int(n ** 0.5) + 1, 2):
        if n % i == 0:
            return False

    # If none of the above conditions are met, the number is prime
    return True

# create a function to do 5 unit tests of the code above
# def test_is_prime():
#     assert is_prime(1) == False
#     assert is_prime(2) == True
#     assert is_prime(3) == True
#     assert is_prime(4) == False
#     assert is_prime(5) == True


# q: what does the function above do?
# a: checks if a number is prime

# explain the code above line-by-line
# 1. Define a function called is_prime that takes a single parameter n
# 2. Check if n is less than 2, if so return False
# 3. Check if n is equal to 2, if so return True
# 4. Check if n is divisible by 2, if so return False
# 5. Iterate over odd numbers from 3 to the square root of n
# 6. Check if n is divisible by i, if so return False
# 7. If none of the above conditions
#    are met, return True

