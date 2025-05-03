namespace TelehealthService.Core.Abstractions;

public interface IAiAssistantService
{
    Task<string> GetAccessTokenAsync();
    Task<string> GetResponseAsync(string prompt);
}