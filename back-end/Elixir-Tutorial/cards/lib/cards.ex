defmodule Cards do
  @moduledoc """
    Provides methods for creating and handling a deck of cards
  """

  @doc """
    Returns a list of strings representing a deck of playing cards
  """
  def create_deck do
    values = ["Ace", "Two", "Three", "Four", "Five"]
    suits = ["Spades", "Clubs", "Hearts", "Diamonds"]
  
    for suit <- suits, value <- values do
      "#{value} of #{suit}"
    end
  end

  def shuffle(deck) do 
    Enum.shuffle(deck)
  end

  @doc """
    Determines wether a deck contains a given card

  ## Examples

      iex> deck = Cards.create_deck
      ["Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades",
      "Five of Spades", "Ace of Clubs", "Two of Clubs", "Three of Clubs",
      "Four of Clubs", "Five of Clubs", "Ace of Hearts", "Two of Hearts",
      "Three of Hearts", "Four of Hearts", "Five of Hearts", "Ace of Diamonds",
      "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds"]
      iex> Cards.contains?(deck, "Ace of Spades")
      true
  """
  def contains?(deck, card) do
    Enum.member?(deck, card)
  end

  @doc """
    Divides a deck into a hand and the remainder of the deck.
    The `hand_size` agument indicates how many cards should be in the hand.

  ## Examples

      iex> deck = Cards.create_deck
      iex> {hand, deck} = Cards.deal(deck, 1)
      iex> hand
      ["Ace of Spades"] 
  """
  def deal(deck, hand_size) do
    Enum.split(deck , hand_size) 
  end

  def save(deck, filename) do
    binary = :erlang.term_to_binary(deck)
    File.write(filename, binary)
  end

  def load(filename) do 
    case File.read(filename) do 
      {:ok, binary} -> :erlang.binary_to_term binary 
      {:error, _reason} -> "That file does not exist!"
    end 
  end

  def create_hand(hand_size) do
    Cards.create_deck
    |> Cards.shuffle # passing the deck to next function
    |> Cards.deal(hand_size) # deck is the first param of the new function 
  end

end
