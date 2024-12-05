defmodule Part1 do
  def safe_greater?(list) do
    case list do
      [] -> true
      [_] -> true
      [a, b | tail] when (a > b and abs(a-b) in 1..3) -> safe_greater?([b | tail])
      _ -> false
    end
  end

  def safe_lesser?(list) do
    case list do
      [] -> true
      [_] -> true
      [a, b | tail] when (a < b and abs(a-b) in 1..3) -> safe_lesser?([b | tail])
      _ -> false
    end
  end

  def safe?(list) do
    safe_greater?(list) || safe_lesser?(list)
  end

  def run() do
    File.stream!("./2.in", :line)
      |> Stream.map(&String.trim/1)
      |> Stream.filter(&(byte_size(&1) > 0))
      |> Stream.map(&String.split(&1, " "))
      |> Stream.map(fn list -> Enum.map(list, &String.to_integer/1) end)
      |> Stream.map(fn list -> safe?(list) end)
      |> Enum.to_list()
      |> Enum.reduce(0, fn x, acc -> case x do true -> acc + 1; _ -> acc end end)
      |> IO.inspect()
  end
end

defmodule Part2 do
  def listlists(list) do
    Enum.with_index(list)
    |> Enum.map(fn {_, index} -> List.delete_at(list, index) end)
  end

  def safe?(list) do
    listlists(list)
    |> IO.inspect()
  end

  def run() do
    File.stream!("./2.in", :line)
      |> Stream.map(&String.trim/1)
      |> Stream.filter(&(byte_size(&1) > 0))
      |> Stream.map(&String.split(&1, " "))
      |> Stream.map(fn list -> Enum.map(list, &String.to_integer/1) end)
      |> Stream.map(fn list -> safe?(list) end)
      |> Enum.to_list()
      # |> Enum.reduce(0, fn x, acc -> case x do true -> acc + 1; _ -> acc end end)
      |> IO.inspect()
  end
end

Part1.run()
Part2.run()
