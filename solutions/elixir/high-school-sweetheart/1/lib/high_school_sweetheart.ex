defmodule HighSchoolSweetheart do
  def first_letter(name), do: String.trim(name) |> String.at(0)

  def initial(name) do
    first_letter(name)
    |> String.trim()
    |> String.upcase()
    |> Kernel.<>(".")
  end

  def initials(full_name) do
    String.split(full_name)
    |> Enum.map(fn x -> initial(x) end)
    |> Enum.join(" ")
  end

  def pair(full_name1, full_name2) do
    """
         ******       ******
       **      **   **      **
     **         ** **         **
    **            *            **
    **                         **
    **     #{initials(full_name1)}  +  #{initials(full_name2)}     **
     **                       **
       **                   **
         **               **
           **           **
             **       **
               **   **
                 ***
                  *
    """
  end
end
