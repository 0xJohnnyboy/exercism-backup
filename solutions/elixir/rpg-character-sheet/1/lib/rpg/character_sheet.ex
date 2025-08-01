defmodule RPG.CharacterSheet do

  def welcome() do
    IO.puts("Welcome! Let's fill out your character sheet together.")
    :ok
  end

  defp ask(question) do
    IO.gets(question) |> String.trim()
  end

  def ask_name() do
    ask("What is your character's name?\n")
  end

  def ask_class() do
    ask("What is your character's class?\n")
  end

  def ask_level() do
    ask("What is your character's level?\n") |> String.to_integer()
  end

  def run() do
    welcome()
    %{:name => ask_name(), :class => ask_class(), :level => ask_level()}
    |> IO.inspect(label: "Your character")
  end
end
