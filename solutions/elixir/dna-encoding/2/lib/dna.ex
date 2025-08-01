defmodule DNA do
  @dna_map %{?A => 0b001, ?C => 0b010, ?G => 0b0100, ?T => 0b1000, ?\s => 0b0000}

  def encode_nucleotide(code_point), do: Map.get(@dna_map, code_point)

  def decode_nucleotide(encoded_code) do
    @dna_map 
    |> Enum.map(fn {k, v} -> {v, k} end)
    |> Enum.into(%{})
    |> Map.get(encoded_code)
  end

  def encode(dna), do: do_encode(dna, <<>>)

  defp do_encode([], acc), do: acc

  defp do_encode([h | t], acc),
    do: do_encode(t, <<acc::bitstring, encode_nucleotide(h)::4>>)

  def decode(dna), do: do_decode(dna, ~c"")

  defp do_decode(<<>>, acc), do: acc

  defp do_decode(<<first::4, rest::bitstring>>, acc),
    do: do_decode(rest, acc ++ [decode_nucleotide(first)])
end
