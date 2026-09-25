using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.CompilerServices;

namespace Test
{
    public class Solution
    {
        public int[] AsteroidCollision(int[] asteroids)
        {
            if (asteroids.Length == 0) return [];

            Stack<int> output = new Stack<int>();
            output.Push(asteroids[0]); //Thêm phần tử đầu tiên
            int index = 1;

            while (index < asteroids.Length-1) { 
                if(asteroids[index] > 0 && output.First() >0)
                {
                    output.Push(asteroids[index]);
                    index++;
                }
                else if(output.First() > 0 && asteroids[index] < 0)
                {
                    if(Math.Abs(output.First()) > Math.Abs(asteroids[index]))
                    {
                        index++;
                    }
                    else if (Math.Abs(output.First()) < Math.Abs(asteroids[index])) {
                        output.Pop();
                    }else if (Math.Abs(output.First()) == Math.Abs(asteroids[index]))
                    {
                        output.Pop();
                        index++;
                    }
                }
                else
                {
                    output.Push(asteroids[index]);
                    index++;
                }
            }


            Console.WriteLine("[{0}]", string.Join(", ", output.ToArray()));
            Console.WriteLine(output.First());
            return output.ToArray();
        }

        static void Main(string[] args) {
            Solution ob = new Solution();
            Console.WriteLine("Result: [{0}]",string.Join(", ", ob.AsteroidCollision([2, 5,10, 15, -35, -2, -11])));
        }
    }
}
