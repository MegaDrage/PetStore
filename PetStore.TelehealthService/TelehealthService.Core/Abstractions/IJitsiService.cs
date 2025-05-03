namespace TelehealthService.Core.Abstractions;

public interface IJitsiService
{
    string GenerateRoomUrl(string userIdentifier);
}