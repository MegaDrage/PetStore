using Microsoft.Extensions.Options;
using TelehealthService.Core.Abstractions;

namespace TelehealthService.JitsiIntegration.Services;

public class JitsiService : IJitsiService
{
    private readonly JitsiOptions _options;

    public JitsiService(IOptions<JitsiOptions> options)
    {
        _options = options.Value;
    }

    public string GenerateRoomUrl(string userIdentifier)
    {
        return $"{_options.BaseUrl}/{Guid.NewGuid().ToString("n")!}-{userIdentifier}";
    }
}