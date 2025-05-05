using Microsoft.AspNetCore.SignalR;

namespace TelehealthService.Chat.Hubs;

public class ChatHub : Hub
{
    public async Task JoinRoom(string sessionId)
    {
        await Groups.AddToGroupAsync(Context.ConnectionId, sessionId);
    }

    public async Task SendMessage(string sessionId, string message, string senderId)
    {
        await Clients.Group(sessionId).SendAsync("ReceiveMessage", senderId, message);
    }
}