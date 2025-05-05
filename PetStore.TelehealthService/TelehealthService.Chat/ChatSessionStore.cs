using System.Collections.Concurrent;

namespace TelehealthService.Chat;

public static class ChatSessionStore
{
    public static ConcurrentDictionary<string, ChatSession> Sessions { get; } = new();
}