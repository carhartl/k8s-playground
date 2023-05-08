output "region" {
  value = var.region
}

output "project_id" {
  value = var.project_id
}

output "kubernetes_cluster_name" {
  value = google_container_cluster.primary.name
}

output "kubernetes_cluster_host" {
  value = google_container_cluster.primary.endpoint
}
