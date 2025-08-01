defmodule RPG.CharacterSheet do
  def welcome() do
    IO.puts("Welcome! Let's fill out your character sheet together.")
    :ok
  end

  defp ask(question), do: IO.gets(question) |> String.trim()

  def ask_name(), do: ask("What is your character's name?\n")
  def ask_class(), do: ask("What is your character's class?\n")
  def ask_level(), do: ask("What is your character's level?\n") |> String.to_integer()

  def run() do
    welcome()

    %{:name => ask_name(), :class => ask_class(), :level => ask_level()}
    |> IO.inspect(label: "Your character")
  end
end
