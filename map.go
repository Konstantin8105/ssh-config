package ssh_config

import "fmt"

func Convert(name string) (key SSHKey, err error) {
	switch name {

	case "Host":
		return Host, nil

	case "Match":
		return Match, nil

	case "AddKeysToAgent":
		return AddKeysToAgent, nil

	case "AddressFamily":
		return AddressFamily, nil

	case "BatchMode":
		return BatchMode, nil

	case "BindAddress":
		return BindAddress, nil

	case "CanonicalDomains":
		return CanonicalDomains, nil

	case "CanonicalizeFallbackLocal":
		return CanonicalizeFallbackLocal, nil

	case "CanonicalizeHostname":
		return CanonicalizeHostname, nil

	case "CanonicalizeMaxDots":
		return CanonicalizeMaxDots, nil

	case "CanonicalizePermittedCNAMEs":
		return CanonicalizePermittedCNAMEs, nil

	case "CertificateFile":
		return CertificateFile, nil

	case "ChallengeResponseAuthentication":
		return ChallengeResponseAuthentication, nil

	case "CheckHostIP":
		return CheckHostIP, nil

	case "Cipher":
		return Cipher, nil

	case "Ciphers":
		return Ciphers, nil

	case "ClearAllForwardings":
		return ClearAllForwardings, nil

	case "Compression":
		return Compression, nil

	case "CompressionLevel":
		return CompressionLevel, nil

	case "ConnectionAttempts":
		return ConnectionAttempts, nil

	case "ConnectTimeout":
		return ConnectTimeout, nil

	case "ControlMaster":
		return ControlMaster, nil

	case "ControlPath":
		return ControlPath, nil

	case "ControlPersist":
		return ControlPersist, nil

	case "DynamicForward":
		return DynamicForward, nil

	case "EnableSSHKeysign":
		return EnableSSHKeysign, nil

	case "EscapeChar":
		return EscapeChar, nil

	case "ExitOnForwardFailure":
		return ExitOnForwardFailure, nil

	case "FingerprintHash":
		return FingerprintHash, nil

	case "ForwardAgent":
		return ForwardAgent, nil

	case "ForwardX11":
		return ForwardX11, nil

	case "ForwardX11Timeout":
		return ForwardX11Timeout, nil

	case "ForwardX11Trusted":
		return ForwardX11Trusted, nil

	case "GatewayPorts":
		return GatewayPorts, nil

	case "GlobalKnownHostsFile":
		return GlobalKnownHostsFile, nil

	case "GSSAPIAuthentication":
		return GSSAPIAuthentication, nil

	case "GSSAPIKeyExchange":
		return GSSAPIKeyExchange, nil

	case "GSSAPIClientIdentity":
		return GSSAPIClientIdentity, nil

	case "GSSAPIServerIdentity":
		return GSSAPIServerIdentity, nil

	case "GSSAPIDelegateCredentials":
		return GSSAPIDelegateCredentials, nil

	case "GSSAPIRenewalForcesRekey":
		return GSSAPIRenewalForcesRekey, nil

	case "GSSAPITrustDns":
		return GSSAPITrustDns, nil

	case "HashKnownHosts":
		return HashKnownHosts, nil

	case "HostbasedAuthentication":
		return HostbasedAuthentication, nil

	case "HostbasedKeyTypes":
		return HostbasedKeyTypes, nil

	case "HostKeyAlgorithms":
		return HostKeyAlgorithms, nil

	case "HostKeyAlias":
		return HostKeyAlias, nil

	case "HostName":
		return HostName, nil

	case "IdentitiesOnly":
		return IdentitiesOnly, nil

	case "IdentityFile":
		return IdentityFile, nil

	case "IgnoreUnknown":
		return IgnoreUnknown, nil

	case "IPQoS":
		return IPQoS, nil

	case "KbdInteractiveAuthentication":
		return KbdInteractiveAuthentication, nil

	case "KbdInteractiveDevices":
		return KbdInteractiveDevices, nil

	case "KexAlgorithms":
		return KexAlgorithms, nil

	case "LocalCommand":
		return LocalCommand, nil

	case "LocalForward":
		return LocalForward, nil

	case "LogLevel":
		return LogLevel, nil

	case "MACs":
		return MACs, nil

	case "NoHostAuthenticationForLocalhost":
		return NoHostAuthenticationForLocalhost, nil

	case "NumberOfPasswordPrompts":
		return NumberOfPasswordPrompts, nil

	case "PasswordAuthentication":
		return PasswordAuthentication, nil

	case "PermitLocalCommand":
		return PermitLocalCommand, nil

	case "PKCS11Provider":
		return PKCS11Provider, nil

	case "Port":
		return Port, nil

	case "PreferredAuthentications":
		return PreferredAuthentications, nil

	case "Protocol":
		return Protocol, nil

	case "ProxyCommand":
		return ProxyCommand, nil

	case "ProxyUseFdpass":
		return ProxyUseFdpass, nil

	case "PubkeyAcceptedKeyTypes":
		return PubkeyAcceptedKeyTypes, nil

	case "PubkeyAuthentication":
		return PubkeyAuthentication, nil

	case "RekeyLimit":
		return RekeyLimit, nil

	case "RemoteForward":
		return RemoteForward, nil

	case "RequestTTY":
		return RequestTTY, nil

	case "RevokedHostKeys":
		return RevokedHostKeys, nil

	case "RhostsRSAAuthentication":
		return RhostsRSAAuthentication, nil

	case "RSAAuthentication":
		return RSAAuthentication, nil

	case "SendEnv":
		return SendEnv, nil

	case "ServerAliveCountMax":
		return ServerAliveCountMax, nil

	case "ServerAliveInterval":
		return ServerAliveInterval, nil

	case "StreamLocalBindMask":
		return StreamLocalBindMask, nil

	case "StreamLocalBindUnlink":
		return StreamLocalBindUnlink, nil

	case "StrictHostKeyChecking":
		return StrictHostKeyChecking, nil

	case "TCPKeepAlive":
		return TCPKeepAlive, nil

	case "Tunnel":
		return Tunnel, nil

	case "TunnelDevice":
		return TunnelDevice, nil

	case "UpdateHostKeys":
		return UpdateHostKeys, nil

	case "UsePrivilegedPort":
		return UsePrivilegedPort, nil

	case "User":
		return User, nil

	case "UserKnownHostsFile":
		return UserKnownHostsFile, nil

	case "VerifyHostKeyDNS":
		return VerifyHostKeyDNS, nil

	case "VisualHostKey":
		return VisualHostKey, nil

	case "XAuthLocation":
		return XAuthLocation, nil

	}
	return key, fmt.Errorf("Not valid name of ssh key : %s", name)
}
