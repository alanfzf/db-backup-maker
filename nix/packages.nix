{
  pkgs,
  self,
  system,
}:
let

in
{
  app = pkgs.buildGoModule {
    pname = "db-backup-maker";
    version = "0.1.0";
    src = ../.;
    vendorHash = null;
  };

  docker = pkgs.dockerTools.buildImage {
    name = "db-backup-maker";
    tag = "latest";

    copyToRoot = pkgs.buildEnv {
      name = "docker-db-backup-maker";
      paths = [
        pkgs.coreutils
        pkgs.dockerTools.fakeNss
        pkgs.dockerTools.usrBinEnv
        pkgs.dockerTools.binSh
        pkgs.dockerTools.caCertificates
        # actual app
        pkgs.supercronic
        pkgs.awscli2
        pkgs.mariadb.client
        self.packages.${system}.app
      ];
    };

    config = {
      Cmd = [
        "${pkgs.supercronic}/bin/supercronic"
        "/etc/crontab"
      ];
    };
  };

  default = self.packages.${system}.docker;
}
