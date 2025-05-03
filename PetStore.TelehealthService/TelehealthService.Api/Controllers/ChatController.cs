using System.Collections.Concurrent;
using Microsoft.AspNetCore.Mvc;
using TelehealthService.Chat;

[ApiController]
[Route("[controller]")]
public class ChatController : ControllerBase
{
    private static ConcurrentDictionary<string, ChatSession> _sessions = new();

    [HttpPost("create")]
    public IActionResult CreateSession()
    {
        var sessionId = Guid.NewGuid().ToString();
        var session = new ChatSession { SessionId = sessionId };
        
        if (!ChatSessionStore.Sessions.TryAdd(sessionId, session))
            return BadRequest("Session already exists");
        
        return Ok(new { SessionId = sessionId });
    }

    [HttpDelete("{sessionId}")]
    public IActionResult CloseSession(string sessionId)
    {
        if (ChatSessionStore.Sessions.TryRemove(sessionId, out _))
            return Ok();
        return NotFound();
    }
}