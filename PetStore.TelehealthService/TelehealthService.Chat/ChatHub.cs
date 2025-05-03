using System.Collections.Concurrent;
using Microsoft.AspNetCore.SignalR;
using TelehealthService.Chat;

public class ChatHub : Hub
{
    public async Task JoinChat(string sessionId, string userId)
    {
        if (!ChatSessionStore.Sessions.TryGetValue(sessionId, out var session))
            throw new HubException("Session not found");

        if (!session.IsActive)
            throw new HubException("Session is closed");
        
        if (string.IsNullOrEmpty(session.ClientId))
            session.ClientId = userId;
        else if (string.IsNullOrEmpty(session.DoctorId))
            session.DoctorId = userId;
        else
            throw new HubException("Chat is full");

        await Groups.AddToGroupAsync(Context.ConnectionId, sessionId);
        await Clients.Group(sessionId).SendAsync("UserJoined", userId);
    }

    public async Task SendMessage(string sessionId, string sender, string message)
    {
        await Clients.Group(sessionId).SendAsync("ReceiveMessage", sender, message);
    }

    public async Task CloseChat(string sessionId)
    {
        if (ChatSessionStore.Sessions.TryGetValue(sessionId, out var session))
        {
            session.IsActive = false;
            await Clients.Group(sessionId).SendAsync("ChatClosed");
        }
    }
}

public class ChatSession
{
    public string SessionId { get; set; }
    public string ClientId { get; set; }
    public string DoctorId { get; set; }
    public bool IsActive { get; set; } = true;
}

public class ChatMessage
{
    public string Sender { get; set; }
    public string Content { get; set; }
    public DateTime Timestamp { get; set; }
}