repeat wait() until game:IsLoaded() == true
wait(8)

local players = game:GetService("Players")
local localplayer = players.LocalPlayer
repeat wait() until localplayer.Character and localplayer.Character.HumanoidRootPart
local HumanoidRootPart = localplayer.Character:WaitForChild("HumanoidRootPart")
local ReplicatedStorage = game:GetService("ReplicatedStorage")
local TeleportService = game:GetService("TeleportService")
local VirtualUser = game:GetService("VirtualUser")
local StarterGui = game:GetService("StarterGui")

-- so I decided to not leak the full bot source, (Prison Life) it should be here now but I wont leak this due to it all being broken (like legit 90% of it is broken) so yea bye nuggies
